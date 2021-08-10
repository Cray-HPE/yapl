import * as yaml from "js-yaml";
import { makeObservable, observable } from "mobx";
import { createContext, useContext } from "react";
interface IMetadata {
  name: string;
  description: string;
  renderedId: number;
  parentId: string;
  id: string;
  childrenIds: string[];
  status: string;
}

export interface IYapl {
  kind: string;
  metadata: IMetadata;
  spec: any;
}

export class YaplStore {
  @observable
  public yaplList: IYapl[] = [];
  @observable
  public SelectedObj: IYapl;
  private ws;

  constructor() {
    this.ws = new WebSocket("ws://localhost:8080/ws");
    this.connect();
    this.reload();
    this.SelectedObj = this.yaplList[0]
    makeObservable(this);
  }

  public reload = () => {
    this.yaplList = [];
    let headers = new Headers();
    headers.append("pragma", "no-cache");
    fetch(`/render`, { headers: headers })
      .then((res) => res.json())
      .then((yaplList) => {
        yaplList.forEach((yapl: IYapl)=>{
          this.addYapl(yapl)
        })
        
      });
  };

  public addYapl = (yapl: IYapl) => {
    this.yaplList.push(yapl);
  };

  public resetStatus = ()=>{
    this.yaplList = this.yaplList.map((yapl: IYapl)=>{
      yapl.metadata.status = "Not Started"
      return yapl
    })
  }

  public updateYapl = (updatedYapl: any) => {
    const updatedYapls = this.yaplList.map((yapl) => {
      if (yapl.metadata.id === updatedYapl.metadata.id) {
        return { ...updatedYapl };
      }
      return yapl;
    });
    this.yaplList = updatedYapls;
    this.SelectedObj = updatedYapl
  };
  connect = () => {
    console.log("Attempting Connection...");

    this.ws.onopen = () => {
      console.log("Successfully Connected");
    };

    this.ws.onmessage = (msg) => {
      let obj = yaml.load(msg.data) as IYapl;
      this.updateYapl(obj);
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
    throw new Error("useStores must be used winin a provider");
  }
  return store;
};
