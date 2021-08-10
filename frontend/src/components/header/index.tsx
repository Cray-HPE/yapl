import { CaretRightFilled, StepForwardFilled } from "@ant-design/icons";
import { Badge, Tooltip, Button, Layout } from "antd";
import { Header as AntdHeader } from "antd/lib/layout/layout";
import { useObserver } from "mobx-react-lite";
import { useStores } from "../../stores/YaplStore";
export const Header = () => {
  const { YaplStore } = useStores();

  const startPipeline = () => {
    YaplStore.reload()
    let headers = new Headers()
    headers.append("pragma", "no-cache")
    fetch(`/start`,{headers: headers})
      .then((res) => res.json())
      .then((json) => console.log(json.data));
  };

  return useObserver(() => (
    <>
      <AntdHeader style={{ background: "#01A982" }}>
        <Badge status="success" text="Success" />
        <div style={{ float: "right" }}>
          <Tooltip title="Run from Beginning">
            <Button
              shape="circle"
              icon={<CaretRightFilled />}
              style={{ marginRight: "20px" }}
              onClick={startPipeline}
            />
          </Tooltip>
          <Tooltip title="Resume from last run">
            <Button
              shape="circle"
              icon={<StepForwardFilled />}
              style={{ marginRight: "20px" }}
              onClick={resumePipeline}
            />
          </Tooltip>
        </div>
      </AntdHeader>
    </>
  ));
};

const resumePipeline = () => {};
