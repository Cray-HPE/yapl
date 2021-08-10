import { Col, Tree } from "antd";
import { Observer, useObserver } from "mobx-react-lite";
import { IYapl, useStores } from "../../stores/YaplStore";
import { trace } from "mobx";
import {
  DownOutlined,
  FrownFilled,
  FrownOutlined,
  MehOutlined,
  SmileOutlined,
} from "@ant-design/icons";
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
        onSelect={(key)=>{
            const obj = YaplStore.yaplList.find(yapl=> yapl.metadata.id === key[0])
            YaplStore.SelectedObj = obj || YaplStore.SelectedObj
        }}
      >
        {YaplStore.yaplList.map((yapl: IYapl) => {
          return (
            <TreeNode
              title={yapl?.metadata?.name + "  -----  " + yapl?.metadata?.status}
              key={yapl?.metadata?.id}
            >
              <TreeNode isLeaf={true} title="1"></TreeNode>
            </TreeNode>
          );
        })}
      </Tree>
    </Col>
  ));
};
