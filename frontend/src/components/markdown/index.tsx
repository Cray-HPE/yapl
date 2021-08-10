import { Col } from "antd";
import { useObserver } from "mobx-react-lite";
import ReactMarkdown from "react-markdown";

export const Markdown = () => {
  return useObserver(() => (
    <Col
      flex="auto"
      style={{ margin: "0", height: "80vh", overflow: "auto" }}
      className="site-layout-background"
    >
      <ReactMarkdown>
        {`# Hello, *world*!
    
## sadfsadf
## sadfsadf    
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf
## sadfsadf

`}
      </ReactMarkdown>
    </Col>
  ));
};
