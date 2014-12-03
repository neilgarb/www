(function($) {

    //--------------------------------------------------------------------------
    // Helper functions
    //--------------------------------------------------------------------------

    function debug() {
        $.each(arguments, function(k, v) {
            console.log(v);
        });
    }

    //--------------------------------------------------------------------------
    // Flash
    //--------------------------------------------------------------------------

    var Flash = {};

    Flash.info = function(msg) {
        debug('Flash::info');
        var $p = $('<p/>')
            .addClass('alert alert-success')
            .html(msg)
            .prependTo($('#flash'));
        setTimeout(function() { $p.fadeOut(); }, 4000);
    };

    Flash.error = function(msg) {
        debug('Flash::error');
        var $p = $('<p/>')
            .addClass('alert alert-danger')
            .html(msg)
            .prependTo($('#flash'));
        setTimeout(function() { $p.fadeOut(); }, 4000);
    };

    //--------------------------------------------------------------------------
    // A
    //--------------------------------------------------------------------------

    var A = function($a) {
        this.$a = $a;
        if (/^\//.test(this.$a.attr('href'))) {
            this.$a.click(this.onClick.bind(this));
        }
    };

    A.prototype.onClick = function(e) {
        debug('A::onClick');
        e.preventDefault();
        History.pushState(null, this.$a.attr('title'), this.$a.attr('href'));
    };

    //--------------------------------------------------------------------------
    // Form
    //--------------------------------------------------------------------------

    var Form = function($form) {
        this.$form = $form;
        this.$form.submit(this.onSubmit.bind(this));
    };

    Form.prototype.onSubmit = function(e) {
        debug('Form::onSubmit');
        e.preventDefault();
        this.disableSubmitButtons();
        $.ajax({
            url: this.$form.attr('action'),
            type: this.$form.attr('method'),
            success: this.onSuccess.bind(this),
            error: this.onError.bind(this)
        });
    };

    Form.prototype.disableSubmitButtons = function() {
        debug('Form::disableSubmitButtons');
        this.$form.find('button[type=submit]').each(function() {
            var $btn = $(this);
            $btn.attr('data-label', $btn.html());
            $btn.html('Loading...');
            $btn.attr('disabled', 'disabled');
        });
    };

    Form.prototype.enableSubmitButtons = function() {
        debug('Form::enableSubmitButtons');
        this.$form.find('button[type=submit]').each(function() {
            var $btn = $(this);
            $btn.html($btn.attr('data-label'));
            $btn.removeAttr('disabled');
        });
    };

    Form.prototype.onSuccess = function() {
        debug('Form::onSuccess');
        this.enableSubmitButtons();
        Flash.info('Success!');
    };

    Form.prototype.onError = function() {
        debug('Form::onError');
        this.enableSubmitButtons();
        Flash.info('Error!');
    };

    //--------------------------------------------------------------------------
    // Init
    //--------------------------------------------------------------------------

    (function() {
        $('form').each(function() { new Form($(this)); });
        $('a').each(function() { new A($(this)); });
        History.Adapter.bind(window, 'statechange', function() {
            debug('History::statechange');
            var State = History.getState();
            debug(State);
        });
    })();

})(jQuery);
