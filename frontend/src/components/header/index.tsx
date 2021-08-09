import { CaretRightFilled, StepForwardFilled } from "@ant-design/icons";
import { Badge, Tooltip, Button, Layout } from "antd";
import { Header as AntdHeader } from "antd/lib/layout/layout";
import { useObserver } from "mobx-react-lite";
export const Header = () => {
  return useObserver(()=>(
    <>
      <AntdHeader style={{ background: "#01A982" }}>
        <Badge status="success" text="Success" />
        <div style={{ float: "right" }}>
          <Tooltip title="Run from Beginning">
            <Button
              shape="circle"
              icon={<CaretRightFilled />}
              style={{ marginRight: "20px" }}
            />
          </Tooltip>
          <Tooltip title="Resume from last run">
            <Button
              shape="circle"
              icon={<StepForwardFilled />}
              style={{ marginRight: "20px" }}
            />
          </Tooltip>
        </div>
      </AntdHeader>
    </>
  ));
};
