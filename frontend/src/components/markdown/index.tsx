import { Col } from "antd";
import { useObserver } from "mobx-react-lite";
import ReactMarkdown from "react-markdown";
import { useStores } from "../../stores/YaplStore";

export const Markdown = () => {
  const { YaplStore } = useStores();
  return useObserver(() => (
    <Col
      flex="20vw"
      style={{ margin: "0", height: "80vh", overflow: "auto" }}
      className="site-layout-background"
    >
      <div style={{padding:"16px"}}>
        <ReactMarkdown>
          {YaplStore?.SelectedObj?.metadata?.description || ""}
        </ReactMarkdown>
      </div>
    </Col>
  ));
};
