import "./styles/Table.scss";

function Table({
  tableName,
  tableDescription,
  totalCashChange,
  createdBy,
  createdAt,
}) {
  return (
    <>
      <div className="table-container">
        <div className="title-and-description">
          <h1>{tableName}</h1>
          <p>{tableDescription}</p>
        </div>
        <div id="table-info">
          <div id="total-cash-change">{totalCashChange}</div>
          <div id="created-by">{createdBy}</div>
          <div id="created-at">{createdAt}</div>
        </div>
      </div>
    </>
  );
}

export default Table;
