import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";
import { Badge, Button, Col, Layout, Menu, Progress, Row, Tooltip } from "antd";
import ReactMarkdown from "react-markdown";
import { Header } from "./components/header";
import { Markdown } from "./components/markdown";
import { ProgressPage } from "./components/progress";

const { SubMenu } = Menu;
const { Content, Footer, Sider } = Layout;
class App extends Component {
  render() {
    return (
      <Layout style={{ height: "100vh" }}>
        <Header />
        <Content style={{ padding: "5vh 2vw" }}>
          <Layout style={{ padding: "0", height: "100%" }}>
            <Row style={{ marginLeft: 0 }} className="site-layout-background">
              <ProgressPage />

              <Markdown />
            </Row>
          </Layout>
        </Content>
      </Layout>
    );
  }
}

export default App;
