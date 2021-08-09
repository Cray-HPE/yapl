import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";
import { connect } from "./api/ws";
import { Button, Col, Layout, Menu, Row } from "antd";
import {
  UserOutlined,
  LaptopOutlined,
  NotificationOutlined,
} from "@ant-design/icons";
import ReactMarkdown from "react-markdown";

const { SubMenu } = Menu;
const { Header, Content, Footer, Sider } = Layout;
class App extends Component {
  constructor(props: any) {
    super(props);
    connect();
  }
  render() {
    return (
      <Layout style={{ height: "100vh" }}>
        <Header style={{ background: "#01A982" }}>
          <Button type="primary">Primary Button</Button>
        </Header>
        <Content style={{ padding: "5vh 2vw" }}>
          <Layout style={{ padding: "0", height: "100%" }}>
            <Row style={{ marginLeft: 0 }} className="site-layout-background">
              <Col
                flex={3}
                style={{
                  background: "#24292f",
                  height: "80vh",
                  overflow: "auto",
                  color: "white",
                  margin: "0",
                }}
              >
                <div>
                  ...
                  <br />
                  Really
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  long
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  ...
                  <br />
                  content
                </div>
              </Col>
              <Col
                flex={2}
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
            </Row>
          </Layout>
        </Content>
      </Layout>
    );
  }
}

export default App;
