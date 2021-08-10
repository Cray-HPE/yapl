import {
  CheckCircleTwoTone,
  CloseCircleTwoTone,
  DownOutlined,
  QuestionCircleTwoTone,
  SyncOutlined,
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
        overflowX: "hidden",
        color: "white",
        margin: "0",
      }}
    >
      <Tree
        showIcon
        defaultExpandAll
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
              title={<strong>{yapl?.metadata?.name}</strong>}
              key={yapl?.metadata?.id}
              icon={
                yapl?.kind === "step" &&
                (yapl?.metadata?.status === "Not Started" ? (
                  <QuestionCircleTwoTone />
                ) : yapl?.metadata?.status === "Running" ? (
                  <SyncOutlined spin twoToneColor="blue" />
                ) : yapl?.metadata?.status === "Done" ? (
                  <CheckCircleTwoTone twoToneColor="#52c41a" />
                ) : (
                  <CloseCircleTwoTone twoToneColor="red" />
                ))
              }
            >
              {yapl?.kind === "step" && (
                <>
                  <TreeNode
                    isLeaf={true}
                    title="Pre condition"
                    key={yapl?.metadata?.id + "Pre Condition"}
                  />
                  <TreeNode
                    isLeaf={true}
                    title="Action"
                    key={yapl?.metadata?.id + "Action"}
                  />
                  <TreeNode
                    isLeaf={true}
                    title="Post Validation"
                    key={yapl?.metadata?.id + "Post Validation"}
                  />
                </>
              )}
            </TreeNode>
          );
        })}
      </Tree>
    </Col>
  ));
};
