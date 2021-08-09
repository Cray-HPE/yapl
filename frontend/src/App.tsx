import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";
import { connect } from "./api/ws";
import { Col, Layout, Menu, Row } from "antd";
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
        <Header style={{ background: "#01A982", height: "8vh" }}>Header</Header>
        <Content style={{ padding: "5vh 2vw" }}>
          <Layout
            className="site-layout-background"
            style={{ padding: "0", height: "100%" }}
          >
            <Row style={{ margin: 0 }}>
              <Col
                flex={3}
                style={{
                  background: "#24292f",
                  height: "82vh",
                  overflow: "auto",
                  color: "white"
                }}
              >
                <div
                  style={{ margin: "1vh", height: "80vh", overflow: "auto"  }}
                >
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
                style={{ margin: "1vh", height: "80vh", overflow: "auto" }}
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
