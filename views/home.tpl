{{define "body"}}
<div class="home">
    <div class="container">
        <h1>Neil Garb <small>Web and Software Development</small></h1>
        <div class="row">
            <div class="column" style="width: 33.333333%;">
                <h2>Who am I?</h2>
                <p>My name is Neil Garb and I'm a software engineer based in Cape Town, South Africa. My specialities are PHP and Python web development, though I build software in Java and Go as well. I have a Bachelor of Science degree from the University of Cape Town, and have been building software for 15 years.</p>
            </div>
            <div class="column" style="width: 33.333333%;">
                <h2>What do I offer?</h2>
                <p>I am available for short-term projects on a freelance basis, or for longer-term projects on a contract or retainer basis. For your money you'll get quality, robust, hand-crafted code. Anything less I won't charge you for.</p>
            </div>
            <div class="column" style="width: 33.333333%;">
                <h2>What do I charge?</h2>
                <p>This depends largely on the duration of the project and on the type of work. I normally charge an hourly rate, but am negotiable on larger projects.</p>
            </div>
        </div>
        <h2>What work have I done?</h2>
        <div class="row">
            <div class="column" style="width: 66.666666%;">
                <div class="work-list">
                    {{range .Works}}
                    {{template "work" .}}
                    {{end}}
                </div>
            </div>
            <div class="column" style="width: 33.333333%;">
                <dl class="tech-list">
                    {{range $k, $v := .Techs}}
                    <dt>{{$k}}</dt>
                        <dd>
                            <span style="width: {{$v}}%;"></span>
                        </dd>
                    {{end}}
                </dl>
            </div>
        </div>
        <h2>What budgets do I work with?</h2>
        <div class="row">
            <div class="column" style="width: 33.333333%;">
                <h3>$</h3>
                <p>You need a simple website with little or no use of a database, or a piece of software with a simple, self-contained function.</p>
                <p><strong>Duration:</strong> days</p>
            </div>
            <div class="column" style="width: 33.333333%;">
                <h3>$$</h3>
                <p>You need a more complex website or piece of software, with active use of a database or with more complex integrations with third party software such as search engines or queues.</p>
                <p><strong>Duration:</strong> weeks</p>
            </div>
            <div class="column" style="width: 33.333333%;">
                <h3>$$$</h3>
                <p>You need a full bespoke sotware solution with complex inner workings and integrations with third-party software. You'll require regular milestone deliverables and close contact over the course of the project.</p>
                <p><strong>Duration:</strong> months</p>
            </div>
        </div>
        <h2>How can you get in touch with me?</h2>
        <div class="row">
            <div class="column" style="width: 20%;">
                <a class="contact contact-email" href="mailto:info@neilgarb.co.za">
                    <span></span>
                    info@neilgarb.co.za
                </a>
            </div>
            <div class="column" style="width: 20%;">
                <a class="contact contact-twitter" href="https://twitter.com/neil_garb">
                    <span></span>
                    @neil_garb
                </a>
            </div>
            <div class="column" style="width: 20%;">
                <a class="contact contact-skype" href="skype:NeilGarb">
                    <span></span>
                    NeilGarb
                </a>
            </div>
            <div class="column" style="width: 20%;">
                <a class="contact contact-linkedin" href="">
                    <span></span>
                    Neil Garb
                </a>
            </div>
            <div class="column" style="width: 20%;">
                <a class="contact contact-github" href="https://github.com/NeilGarb">
                    <span></span>
                    NeilGarb
                </a>
            </div>
        </div>
    </div>
</div>
{{end}}

{{define "work"}}
<div class="work" data-type="{{.Type}}" data-techs=" {{range .Techs}}{{.Slug}} {{end}}">
    <div class="work-body">
        <h4>{{.Name}}</h4>
        <p>{{.Description}}</p>
    </div>
</div>
{{end}}
