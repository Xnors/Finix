import "./styles/Home.scss";
import "sober";
import Table from "./Table";

function Home() {
  return (
    <>
      <div className="top-bar">
        <div className="logo-box">Finix </div>
        <SoberInputBox />
      </div>
      <div className="content">
        <s-button type="elevated"> New Table</s-button>
        <div className="tables-box"> 
          <Table
            tableName={"Table 1"}
            tableDescription={"Lorem ipsum dolor sit amet consectetur, adipisicing elit. Beatae ipsum voluptates dolor amet mollitia aliquam quos molestiae soluta doloremque debitis rem numquam, voluptatum, facere consequuntur, hic possimus sint! Voluptate, similique!"}
            totalCashChange={"+114514"}
            createdBy={"Fexcode"}
            createdAt={"2021-01-01"}
          />
          <Table
            tableName={"Xnors 专用"}
            tableDescription={"Lorem ipsum dolor sit amet consectetur, adipisicing elit. Beatae ipsum voluptates dolor amet mollitia aliquam quos molestiae soluta doloremque debitis rem numquam, voluptatum, facere consequuntur, hic possimus sint! Voluptate, similique!"}
            totalCashChange={"+3"}
            createdBy={"XnorsCode"}
            createdAt={"2025-03-22"}
          />
        </div>
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
