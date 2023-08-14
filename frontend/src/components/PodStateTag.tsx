import React, { memo } from "react";
import pb from "../api/compiled";
import { Tag } from "antd";
import { SyncOutlined, LoadingOutlined } from "@ant-design/icons";
import { css } from "@emotion/css";
import theme from "../styles/theme";
const PodStateTag: React.FC<{ pod: pb.types.StateContainer }> = ({ pod }) => {
  if (pod.is_old) {
    return (
      <Tag
        icon={<SyncOutlined spin />}
        color="#fde047"
        style={{ marginLeft: 5 }}
      >
        {pod.pod} 即将停止
      </Tag>
    );
  }

  if (pod.terminating) {
    return (
      <Tag
        icon={<SyncOutlined spin />}
        color="#fca5a5"
        style={{ marginLeft: 5 }}
      >
        {pod.pod} 停止中
      </Tag>
    );
  }

  if (pod.pending) {
    return (
      <Tag
        icon={<LoadingOutlined spin />}
        color="#67e8f9"
        style={{ marginLeft: 5 }}
      >
        {pod.pod} 未就绪
      </Tag>
    );
  }

  if (!pod.is_old && !pod.ready) {
    return (
      <Tag
        icon={<SyncOutlined spin />}
        color="#93c5fd"
        style={{ marginLeft: 5 }}
      >
        {pod.pod} 启动中
      </Tag>
    );
  }

  return (
    <Tag
      className={css`
        color: white;
        background-color: ${theme.lightMainColor};
        border-color: transparent;
      `}
      style={{ marginLeft: 5 }}
    >
      {pod.pod}
    </Tag>
  );
};

export default memo(PodStateTag);
