import "./styles/Home.scss";
import "sober";
import Table from "./Table";
import { createSignal, createEffect, For, createResource } from "solid-js";

function Home() {
  const [data, setData] = createSignal(null);
  createEffect(() => {
    fetch("/get_info")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((json) => {
        setData(json);
        console.log("获取数据成功! ", JSON.stringify(data()));
      })
      .catch((error) => console.error(error));
  }, []);

  let tables_elm = (
    <div className="tables-box">
      <For each={data()}>
        {(table) => (
          <Table
            tableName={table.name.substring(0, table.name.length - 5)} // 去掉后缀: .json
            tableDescription={table.table_info.description}
            totalCashChange={"+3"}
            createdBy={"XnorsCode"}
            createdAt={table.table_info.created_at.split(" ")[0]} // 日期只取年月日
          />
        )}
      </For>
    </div>
  );

  return (
    <>
      <div className="top-bar">
        <div className="logo-box">Finix </div>
        <SoberInputBox />
      </div>
      <div className="content">
        <s-button type="elevated"> New Table</s-button>
        {tables_elm}
      </div>
    </>
  );
}

function SoberInputBox() {
  return (
    <>
      <s-text-field label="请输入内容" className="input-field">
        <s-icon slot="start" name="search" className="search-icon"></s-icon>
        <s-icon-button slot="end">
          <s-icon name="close"></s-icon>
        </s-icon-button>
      </s-text-field>
    </>
  );
}

export default Home;
