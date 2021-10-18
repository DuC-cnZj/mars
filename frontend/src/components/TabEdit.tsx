import React, { useState, useEffect, useCallback, memo } from "react";
import { Controlled as CodeMirror } from "react-codemirror2";
import PipelineInfo from "./PipelineInfo";
import { commit, configFile, projects } from "../api/gitlab";
import {
  DeployStatus as DeployStatusEnum,
  selectList,
} from "../store/reducers/createProject";
import { Button, Skeleton, Progress, message } from "antd";
import "codemirror/lib/codemirror.css";
import "codemirror/theme/material.css";
import "codemirror/theme/dracula.css";
import { useSelector, useDispatch } from "react-redux";
import {
  clearCreateProjectLog,
  setCreateProjectLoading,
  setDeployStatus,
} from "../store/actions";
import { toSlug } from "../utils/slug";
import { useWs, useWsReady } from "../contexts/useWebsocket";
import {
  ArrowLeftOutlined,
  StopOutlined,
  ArrowRightOutlined,
} from "@ant-design/icons";
import classNames from "classnames";
import LogOutput from "./LogOutput";
import ProjectSelector from "./ProjectSelector";
import TimeCost from "./TimeCost";
import DebugModeSwitch from "./DebugModeSwitch";
import pb from "../api/compiled";

require("codemirror/mode/go/go");
require("codemirror/mode/css/css");
require("codemirror/mode/javascript/javascript");
require("codemirror/mode/yaml/yaml");
require("codemirror/mode/php/php");
require("codemirror/mode/textile/textile");

interface CreateItemInterface {
  gitlabProjectId: number;
  gitlabBranch: string;
  gitlabCommit: string;

  name: string;
  config: string;
  debug: boolean;
}

const ModalSub: React.FC<{
  detail: pb.ProjectShowResponse;
  onSuccess: () => void;
}> = ({ detail, onSuccess }) => {
  let id = detail.id;
  let namespaceId = detail.namespace?.id;
  const ws = useWs();
  const wsReady = useWsReady();

  const [editVisible, setEditVisible] = useState<boolean>(true);
  const [timelineVisible, setTimelineVisible] = useState<boolean>(false);
  const list = useSelector(selectList);
  const dispatch = useDispatch();
  const [data, setData] = useState<CreateItemInterface>({
    name: detail.name,
    gitlabProjectId: Number(detail.gitlab_project_id),
    gitlabBranch: detail.gitlab_branch,
    gitlabCommit: detail.gitlab_commit,
    config: detail.config,
    debug: !detail.atomic,
  });
  const [mode, setMode] = useState<string>("text/x-yaml");
  const [initValue, setInitValue] = useState<{
    projectName: string;
    gitlabProjectId: string;
    gitlabBranch: string;
    gitlabCommit: string;
    time?: number;
  }>();
  console.log("namespaceId, data.name", namespaceId, data.name)
  let slug = toSlug(namespaceId, data.name);
  console.log("slug: ", slug)

  // 初始化，设置 initvalue
  useEffect(() => {
    projects().then((res) => {
      if (
        detail &&
        detail.gitlab_project_id &&
        detail.gitlab_branch &&
        detail.gitlab_commit
      ) {
        commit({
          project_id: detail.gitlab_project_id,
          branch: detail.gitlab_branch,
          commit: detail.gitlab_commit,
        }).then((res) => {
          setInitValue({
            projectName: detail.name,
            gitlabProjectId: detail.gitlab_project_id,
            gitlabBranch: detail.gitlab_branch,
            gitlabCommit: res.data.data.label,
          });
        });
      }
    });
  }, [setInitValue, detail]);

  // 更新成功，触发 onSuccess
  useEffect(() => {
    if (list[slug]?.deployStatus === DeployStatusEnum.DeployUpdateSuccess) {
      setStart(false);
      setTimelineVisible(false);
      setEditVisible(true);
      dispatch(setDeployStatus(slug, DeployStatusEnum.DeployUnknown));
      onSuccess();
    }
  }, [list, dispatch, slug, onSuccess]);

  // 更新 config 文件的类型， TODO 支持动态加载 mode css 文件
  const loadConfigFile = useCallback(() => {
    configFile({project_id: String(data.gitlabProjectId), branch: data.gitlabBranch}).then((res) => {
      setData((d) => ({ ...d, config: res.data.data }));
      switch (res.data.type) {
        case "dotenv":
        case "env":
        case ".env":
          setMode("text/x-textile");
          break;
        case "yaml":
          setMode("text/x-yaml");
          break;
        case "php":
          setMode("php");
          break;
        default:
          setMode(res.data.type);
          break;
      }
    });
  }, [data.gitlabProjectId, data.gitlabBranch]);

  const onChange = ({
    projectName,
    gitlabProjectId,
    gitlabBranch,
    gitlabCommit,
  }: {
    projectName: string;
    gitlabProjectId: number;
    gitlabBranch: string;
    gitlabCommit: string;
  }) => {
    setData((d) => ({
      ...d,
      name: projectName,
      gitlabProjectId: gitlabProjectId,
      gitlabBranch: gitlabBranch,
      gitlabCommit: gitlabCommit,
    }));

    if (gitlabCommit !== "" && data.config === "") {
      loadConfigFile();
    }
  };
  useEffect(() => {
    if (!wsReady) {
      setStart(false);
      dispatch(setCreateProjectLoading(slug, false));
    }
  }, [wsReady]);
  const updateDeploy = () => {
    if (!wsReady) {
      message.error("连接断开了");
      return;
    }
    if (data.gitlabCommit && data.gitlabBranch) {
      setStart(true);
      setEditVisible(false);
      setTimelineVisible(true);

      let re = {
        type: "update_project",
        data: JSON.stringify({
          project_id: Number(id),
          gitlab_branch: data.gitlabBranch,
          gitlab_commit: data.gitlabCommit,
          config: data.config,
          atomic: !data.debug,
        }),
      };
      let s = JSON.stringify(re);
      dispatch(setDeployStatus(slug, DeployStatusEnum.DeployUnknown));

      dispatch(clearCreateProjectLog(slug));
      dispatch(setCreateProjectLoading(slug, true));
      ws?.send(s);
    }
  };
  const [start, setStart] = useState(false);

  useEffect(() => {
    if (list[slug]?.deployStatus !== DeployStatusEnum.DeployUnknown) {
      setStart(false);
    }
  }, [list, slug]);

  const onReset = () => {
    setData({
      name: detail.name,
      gitlabProjectId: Number(detail.gitlab_project_id),
      gitlabBranch: detail.gitlab_branch,
      gitlabCommit: detail.gitlab_commit,
      config: detail.config,
      debug: !detail.atomic,
    });
    if (initValue) {
      setInitValue({ ...initValue, time: new Date().getUTCSeconds() });
    }
  };

  const onRemove = useCallback(() => {
    if (!wsReady) {
      message.error("连接断开了");
      return;
    }
    if (data.gitlabProjectId && data.gitlabBranch && data.gitlabCommit) {
      let re = {
        type: "cancel_project",
        data: JSON.stringify({
          namespace_id: Number(namespaceId),
          name: data.name,
        }),
      };

      let s = JSON.stringify(re);
      ws?.send(s);
      return;
    }
  }, [data, ws, namespaceId, wsReady]);

  return (
    <div className="edit-project">
      <PipelineInfo
        projectId={data.gitlabProjectId}
        branch={data.gitlabBranch}
        commit={data.gitlabCommit}
      />
      <div className={classNames({ "display-none": !editVisible })}>
        <div style={{ width: "100%" }}>
          {list[slug]?.output?.length > 0 ? (
            <Button
              style={{ marginBottom: 20 }}
              type="dashed"
              disabled={list[slug]?.isLoading}
              onClick={() => {
                setEditVisible(false);
                setTimelineVisible(true);
              }}
              icon={<ArrowRightOutlined />}
            />
          ) : (
            ""
          )}
          {initValue ? (
            <ProjectSelector value={initValue} onChange={onChange} />
          ) : (
            <Skeleton.Input active style={{ width: 800 }} size="small" />
          )}
        </div>
        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
            paddingBottom: 10,
          }}
        >
          <span>配置文件:</span>
          <DebugModeSwitch
            value={data.debug}
            onchange={(checked: boolean, event: MouseEvent) => {
              setData((data) => ({ ...data, debug: checked }));
            }}
          />
        </div>
        <div style={{ minWidth: 200, marginBottom: 20 }}>
          <CodeMirror
            value={data.config}
            options={{
              mode: mode,
              theme: "dracula",
              lineNumbers: true,
            }}
            onBeforeChange={(editor, d, value) => {
              console.log(editor, d, value);
              setData({ ...data, config: value });
            }}
          />
        </div>
        <div
          className={classNames("edit-project__footer", {
            "edit-project--hidden": list[slug]?.isLoading,
          })}
        >
          <Button
            style={{ marginRight: 5 }}
            disabled={list[slug]?.isLoading}
            onClick={onReset}
          >
            重置
          </Button>
          <Button
            type="primary"
            loading={list[slug]?.isLoading}
            onClick={updateDeploy}
          >
            部署
          </Button>
        </div>
      </div>
      <div
        id="preview"
        style={{ height: "100%", overflow: "auto" }}
        className={classNames("preview", {
          "display-none": !timelineVisible,
        })}
      >
        <div
          style={{ display: "flex", alignItems: "center", marginBottom: 20 }}
        >
          <Button
            type="dashed"
            disabled={list[slug]?.isLoading}
            onClick={() => {
              setEditVisible(true);
              setTimelineVisible(false);
            }}
            icon={<ArrowLeftOutlined />}
          />
          <Progress
            style={{ padding: "0 10px" }}
            percent={list[slug]?.processPercent}
            status="active"
          />
        </div>
        <div
          style={{ display: "flex", alignItems: "center", marginBottom: 10 }}
        >
          <TimeCost start={start} />

          <Button
            type="primary"
            loading={list[slug]?.isLoading}
            onClick={updateDeploy}
            style={{ marginRight: 10, marginLeft: 10 }}
          >
            部署
          </Button>
          <Button
            hidden={
              list[slug]?.deployStatus === DeployStatusEnum.DeployCanceled
            }
            danger
            icon={<StopOutlined />}
            type="dashed"
            onClick={onRemove}
          >
            取消
          </Button>
        </div>
        <LogOutput slug={slug} />
      </div>
    </div>
  );
};

export default memo(ModalSub);
