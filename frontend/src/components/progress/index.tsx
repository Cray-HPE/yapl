import {
    CheckCircleOutlined,
    CloseCircleOutlined,
    DownOutlined, MinusCircleOutlined, SyncOutlined
} from "@ant-design/icons";
import { Col, Tree } from "antd";
import { useObserver } from "mobx-react-lite";
import { IYapl, useStores } from "../../stores/YaplStore";
import "./index.css";
const { TreeNode } = Tree;
export const ProgressPage = () => {
  const { YaplStore } = useStores();

  return useObserver(() => (
    <Col
      flex="40vw"
      style={{
        background: "#24292f",
        height: "80vh",
        overflow: "auto",
        color: "white",
        margin: "0",
      }}
    >
      <Tree
        showIcon
        defaultExpandAll
        defaultSelectedKeys={["0-0-0"]}
        selectedKeys={[YaplStore?.SelectedObj?.metadata?.id]}
        switcherIcon={<DownOutlined />}
        multiple={false}
        onSelect={(key) => {
          const obj = YaplStore.yaplList.find(
            (yapl) => yapl.metadata.id === key[0]
          );
          YaplStore.SelectedObj = obj || YaplStore.SelectedObj;
        }}
      >
        {YaplStore.yaplList.map((yapl: IYapl) => {
          return (
            <TreeNode
              title={yapl?.metadata?.name}
              key={yapl?.metadata?.id}
              icon={
                yapl?.metadata?.status === "Not Started" ? (
                  <MinusCircleOutlined />
                ) : yapl?.metadata?.status === "Running" ? (
                  <SyncOutlined spin />
                ) : yapl?.metadata?.status === "Done" ? (
                  <CheckCircleOutlined />
                ) : (
                  <CloseCircleOutlined />
                )
              }
            >
              <TreeNode isLeaf={true} title="1"></TreeNode>
            </TreeNode>
          );
        })}
      </Tree>
    </Col>
  ));
};
