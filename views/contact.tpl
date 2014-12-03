{{define "body"}}
<h1>Contact</h1>
<div class="row">
    <div class="col-md-6">
    <form action="/contact" method="post" class="well">
        <div class="form-group">
            <input type="text" name="name" placeholder="Your name" class="form-control" />
        </div>
        <div class="form-group">
            <input type="text" name="email" placeholder="Your email address" class="form-control" />
        </div>
        <div class="form-group">
            <textarea name="message" placeholder="Your message" class="form-control" rows="10"></textarea>
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
</div>
{{end}}
