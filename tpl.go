package main

var tpl = []byte(`<!DOCTYPE html>
<html>
<head>
	<title>Golang-Mongo-Sample</title>
</head>
<body>
	<h1 align="center">It works!</h1>
	<hr />
	<table border="1" align="center" width="50%">
		<caption>Contact List</caption>
		<thead>
			<tr>
			<th bgColor=#e5ecf9>Name</th>
			<th bgColor=#e5ecf9>Phone</th>
			</tr>
		</thead>
		<tbody>
		{{with .List}}
		{{range .}}
			<tr>
			<th>{{.Name}}</th>
			<th>{{.Phone}}</th>
			</tr>
		{{end}}
		{{end}}
		</tbody>
	</table>
	<br />
	<form method="post" action="/new">
		<fieldset>
    		<legend>New contact</legend>
			<label for="name">Name: </label>
			<input id="name" type="text" name="name" required/>
			<label for="phone">Phone: </label>
			<input id="phone" type="text" name="phone" required/>
			<input type="submit" />
		</fieldset>
	</form>
	<br />
	<form method="post" action="/drop">
		<input type="submit" value="clear all contact" />
	</form>
</body>
</html>`)
