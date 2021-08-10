import { observable, computed, action, makeObservable } from "mobx";
import { createContext, useContext } from "react";

interface IMetadata {
  Name: string;
  Description: string;
  RenderedId: number;
  ParentId: string;
  Id: string;
  ChildrenIds: string[];
  Status: boolean;
}

export interface IYapl {
  kind: string;
  Metadata: IMetadata;
  Spec: any;
}

export class YaplStore {
  
  @observable
  public yaplList: IYapl[] = [];
  private ws;

  constructor() {
    this.ws = new WebSocket("ws://localhost:8080/ws");
    this.connect();
    makeObservable(this)
  }
  public reload = () => {
    this.yaplList = []
    //TODO: call api to get list
  }

  public addYapl = (yapl: IYapl) => {
    this.yaplList.push(yapl);
  };

  public updateYapl = (updatedYapl: IYapl) => {
    const updatedYapls = this.yaplList.map((yapl) => {
      if (yapl.Metadata.Id === updatedYapl.Metadata.Id) {
        return { ...updatedYapl };
      }
      return yapl;
    });
    this.yaplList = updatedYapls;
  };
  connect = () => {
    console.log("Attempting Connection...");

    this.ws.onopen = () => {
      console.log("Successfully Connected");
    };

    this.ws.onmessage = (msg) => {
      let obj = JSON.parse(msg.data);
      let yapl = {} as IYapl;
      yapl.Metadata = {
        Id: obj.id,
        Status: obj.status,
        Name: "",
        Description: "",
        ParentId: "",
        RenderedId: 1,
        ChildrenIds: [],
      };
      this.addYapl(yapl)
    };

    this.ws.onclose = (event) => {
      console.log("Socket Closed Connection: ", event);
    };

    this.ws.onerror = (error) => {
      console.log("Socket Error: ", error);
    };
  };
}

export const rootStoreContext = createContext({
  YaplStore: new YaplStore(),
});

export const useStores = () => {
  const store = useContext(rootStoreContext);
  if (!store) {
    throw new Error("useStores must be used winin a provider")
  }
  return store
}
