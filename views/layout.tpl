{{define "layout"}}
<!doctype html>

<html>
    <head>
        <meta charset="utf-8" />
        <link rel="stylesheet" href="/static/bootstrap-3.3.1/css/bootstrap.min.css" type="text/css" />
        <title>{{.Title}}</title>
    </head>
    <body>
        <nav class="navbar navbar-default navbar-inverse" role="navigation">
            <div class="container">
                <ul class="nav navbar-nav">
                    <li><a href="/">Home</a></li>
                    <li><a href="/portfolio">Portfolio</a></li>
                    <li><a href="/contact">Contact</a></li>
                </ul>
            </div>
        </nav>
        <main class="container">
            {{template "body" .}}
        </main>
        <div id="flash"></div>
        <script type="text/javascript" src="/static/jquery-2.1.1.min.js"></script>
        <script type="text/javascript" src="/static/jquery.history.min.js"></script>
        <script type="text/javascript" src="/static/bootstrap-3.3.1/js/bootstrap.min.js"></script>
        <script type="text/javascript" src="/static/www.js"></script>
    </body>
</html>
{{end}}
