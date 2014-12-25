{{define "layout"}}
<!doctype html>
<html>
    <head>
        <meta charset="utf-8" />
        <meta name="description" content="{{.Description}}" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <title>{{.Title}}</title>
        <link rel="stylesheet" href="/static/vendor/reset.css" type="text/css" />
        <link rel="stylesheet" href="/static/www.css" type="text/css" />
        <link rel="shortcut icon" href="/static/favicon.ico" />
    </head>
    <body>
        <script type="text/javascript">
        (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
        (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
        m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
        })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
        ga('create', 'UA-57904516-1', 'auto');
        ga('send', 'pageview');
        </script>
        <main> 
            {{template "body" .}}
        </main>
        <footer>
            <p>The source code for this website is available on <a href="https://github.com/NeilGarb/www">Github</a>.<br />Feel free to use it in any way you like.</p>
        </footer>
        <div id="flash"></div>
        <script type="text/javascript" src="/static/vendor/jquery-2.1.1.min.js"></script>
        <script type="text/javascript" src="/static/vendor/masonry.min.js"></script>
        <script type="text/javascript" src="/static/www.js"></script>
    </body>
</html>
{{end}}
