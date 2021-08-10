import { Layout,  Row } from "antd";
import { Component } from "react";
import "./App.css";
import { Header } from "./components/header";
import { Markdown } from "./components/markdown";
import { ProgressPage } from "./components/progress";

const { Content,  } = Layout;
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
