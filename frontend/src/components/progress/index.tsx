import { Col } from "antd";
import { Observer, useObserver } from "mobx-react-lite";
import { IYapl, useStores } from "../../stores/YaplStore";
import { trace } from "mobx";

export const ProgressPage = () => {
  const { YaplStore } = useStores();

  return useObserver(() => (
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
        {YaplStore.yaplList.map((yapl: IYapl) => {
          console.log("fml");
          return (
            <>
              ID: {yapl.Metadata.Id}
              ----
              Status: {yapl.Metadata.Status}
              <br/>
            </>
          );
        })}
      </div>
    </Col>
  ));
};
