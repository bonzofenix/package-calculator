package app

var indexContent = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PACKAGE CALCULATOR</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
	<div class="container mt-5">
  <div class="container mt-5">
		<h1>Package Calculator</h1>
    <form action="/calculate" method="POST">
			<div id="packSizes" class="input-group mb-3">
        <input id="order" type="text" class="form-control" placeholder="Enter packages separated by ," aria-label="Number input" id="inputNumber" name="inputNumber">
      </div>
      <button type="submit" class="btn btn-primary mb-3">Calculate</button>
    </form>
    <h2>Result</h2>
    <div id="resultBox" class="alert alert-success" style="display: none;"></div>
	</div>
</body>
</html>
`
