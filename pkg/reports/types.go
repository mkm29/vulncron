package reports

type VulnerabilitySummary struct {
	Namespace  string
	Name       string
	Repository string
	Tag        string
	Critical   int
	High       int
	Medium     int
	Low        int
	Unknown    int
	None       int
}

type Summary struct {
	TotalCritical int
	TotalHigh     int
	TotalMedium   int
	TotalLow      int
	TotalUnknown  int
	TotalNone     int
}

var ohtml string = `
<html>
  <head>
    <style>
*{
    box-sizing: border-box;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
}
body{
    font-family: Helvetica;
    -webkit-font-smoothing: antialiased;
    background: rgba( 112, 128, 144, 1);
}
h2{
    text-align: center;
    font-size: 18px;
    text-transform: uppercase;
    letter-spacing: 1px;
    color: white;
    padding: 30px 0;
}
.build-link{
    display: flex;
    justify-content: center;
    text-align: center;
    font-size: 14px;
}
/* Table Styles */
.table-wrapper{
    margin: 10px 70px 70px;
    box-shadow: 0px 35px 50px rgba( 0, 0, 0, 0.2 );
}
.fl-table {
    border-radius: 5px;
    font-size: 12px;
    font-weight: normal;
    border: none;
    border-collapse: collapse;
    width: 100%;
    max-width: 100%;
    white-space: nowrap;
    background-color: white;
}
.fl-table td, .fl-table th {
    text-align: center;
    padding: 8px;
}
.fl-table td {
    border-right: 1px solid #f8f8f8;
    font-size: 12px;
}
.fl-table thead th {
    color: #ffffff;
    background: #4FC3A1;
}
.fl-table thead th:nth-child(odd) {
    color: #ffffff;
    background: #324960;
}
.fl-table tr:nth-child(even) {
    background: #F8F8F8;
}
    </style>
  </head>
  <body>
    <h2>${title}</h2>
    <div class="table-wrapper">
        <table class="fl-table">
            <thead>
            <tr>
                <th>Namespace</th>
                <th>Image</th>
                <th>Tag</th>
                <th>Critical</th>
                <th>High</th>
                <th>Medium</th>
                <th>Low</th>
            </tr>
            </thead>
            <tbody>
`

// var rhtml string = `
// <tr>
//   <td>${namespace}</td>
//   <td>${image}</td>
//   <td>${tag}</td>
//   <td>${critical}</td>
//   <td>${high}</td>
//   <td>${medium}</td>
//   <td>${low}</td>
// </tr>
// `

var chtml string = `
<tbody>
</table>
</div> 
</body>
</html>
`
