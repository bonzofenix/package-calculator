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
		<h1>Package Calculator</h1>
    <form action="/calculate" method="post">
			<div class="form-group">
				<label for="packSizes">Pack sizes:</label>
				<input id="packSizes" value="{{.PackSizes}}" name="packSizes" type="text" class="form-control" placeholder="Enter packages separated by ," aria-label="Number input">
			</div>
			<div class="form-group">
				<label for="order">order:</label>
				<input id="order" value="{{.Order}}" name="order" type="text" class="form-control" placeholder="Enter order size" aria-label="enter order size">
			</div>
      <button type="submit" class="btn btn-primary mb-3">Calculate</button>
    </form>
    <h2>Result</h2>
    <div id="resultBox" class="alert alert-success">
			{{ range $key, $value := .Result }}
				 <li><strong>{{ $key }}</strong>: {{ $value }}</li>
			{{ end }}
		</div>
	</div>
</body>
</html>
`
