import React, { useEffect, useState, useCallback } from "react";
import { DraggableModalProvider } from "../pkg/DraggableModal/DraggableModalProvider";
import ItemCard from "./ItemCard";
import { Empty, Row, Col } from "antd";
import AddNamespace from "./AddNamespace";
import Setting from "./Setting";
import "../pkg/DraggableModal/index.css";
import { listNamespaces } from "../api/namespace";
import { useSelector, useDispatch } from "react-redux";
import { setNamespaceReload } from "../store/actions";
import { selectReload } from "../store/reducers/namespace";
import { useAuth } from "../contexts/auth";

import pb from "../api/compiled";

const AppContent: React.FC = () => {
  const reloadNamespace = useSelector(selectReload);
  const dispatch = useDispatch();
  const [loading, setLoading] = useState(false);
  const [namespaceItems, setNamespaceItems] = useState<pb.NamespaceItem[]>([]);
  const onNamespaceCreated = ({ id, name }: { id: number; name: string }) => {
    console.log(id, name);
    fetchNamespaces();
  };
  const auth = useAuth()
  const fetchNamespaces = useCallback(() => {
    setLoading(true);
    listNamespaces()
      .then((res) => {
        setNamespaceItems(res.data.data);
        setLoading(false);
      })
      .catch((e) => {
        setLoading(false);
      });
  }, []);

  useEffect(() => {
    fetchNamespaces();
  }, [fetchNamespaces]);

  useEffect(() => {
    console.log("reloadNamespace", reloadNamespace);
    if (reloadNamespace) {
      fetchNamespaces();
      dispatch(setNamespaceReload(false));
    }
  }, [reloadNamespace, dispatch, fetchNamespaces]);

  return (
    <DraggableModalProvider>
      <div className="content">
        <AddNamespace onCreated={onNamespaceCreated} />
        {auth.isAdmin() ? <Setting /> : <></>}

        {namespaceItems.length < 1 ? (
          <Empty description={false} imageStyle={{ height: 300 }} />
        ) : (
          <Row gutter={[16, 16]}>
            {namespaceItems.map((item: pb.NamespaceItem) => (
              <Col md={12} lg={8} sm={12} xs={24} key={item.id}>
                <ItemCard
                  loading={loading}
                  item={item}
                  onNamespaceDeleted={() => fetchNamespaces()}
                />
              </Col>
            ))}
          </Row>
        )}
      </div>
    </DraggableModalProvider>
  );
};

export default AppContent;
