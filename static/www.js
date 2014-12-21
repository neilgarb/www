(function($) {

    //-------------------------------------------------------------------------
    // Helper functions
    //-------------------------------------------------------------------------

    function debug() {
        $.each(arguments, function(k, v) {
            console.log(v);
        });
    }

    //-------------------------------------------------------------------------
    // Flash
    //-------------------------------------------------------------------------

    var Flash = {};

    Flash.add = function(cls, msg) {
        var $p = $('<p/>')
            .addClass(cls)
            .html(msg)
            .prependTo($('#flash'));
        this.remove($p);
        $('#flash p:gt(2)').remove();
    };

    Flash.remove = function($el) {
        setTimeout(function() {
            $el.slideUp({
                complete: function() {
                    $(this).remove();
                }
            });
        }, 4000);
    };

    Flash.info = function(msg) {
        debug('Flash::info');
        this.add('flash flash-info', msg)
    };

    Flash.error = function(msg) {
        debug('Flash::error');
        this.add('flash flash-error', msg)
    };

    //-------------------------------------------------------------------------
    // Form
    //-------------------------------------------------------------------------

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
            dataType: 'json',
            data: this.$form.serialize(),
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

    Form.prototype.onSuccess = function(json) {
        debug('Form::onSuccess');
        this.enableSubmitButtons();
        this.clearErrors();
        Flash.info(json.data);
        this.$form[0].reset();
    };

    Form.prototype.onError = function(response) {
        debug('Form::onError');
        this.enableSubmitButtons();
        this.clearErrors();
        var json = JSON.parse(response.responseText);
        switch (json.code) {
            case 400:
                this.setErrors(json.data);
                break;
            default:
                Flash.error(json.data);
                break;
        }
    };

    Form.prototype.clearErrors = function() {
        this.$form.find('.has-error').removeClass('has-error');
        this.$form.find('p.error').remove();
    };

    Form.prototype.setErrors = function(errors) {
        debug('Form::setErrors');
        if (errors) {
            Flash.error('Please fix the errors in your submission.');
            var instance = this;
            $.each(errors, function(k, v) {
                var $formGroup = instance
                    .$form
                    .find('[data-field=' + k + ']');
                $formGroup.addClass('has-error');
                $('<p/>')
                    .addClass('error')
                    .html(v[0])
                    .appendTo($formGroup);
            });
        }
    };

    //-------------------------------------------------------------------------
    // ContactLink
    //-------------------------------------------------------------------------

    var ContactLink = function($el) {
        this.$el = $el;
        this.$el.click(this.onClick.bind(this));
    };

    ContactLink.prototype.onClick = function(e) {
        var href = this.$el.attr('href');
        ga(
            'send',
            'event',
            'outbound',
            'click',
            href,
            {
                'hitCallback': function () {
                    document.location = href;
                }
            }
        );
    };

    //-------------------------------------------------------------------------
    // Init
    //-------------------------------------------------------------------------

    $('form').each(function() { new Form($(this)); });
    $('.work-list').masonry({ columnWidth: 320 });
    $('a.contact').each(function() { new ContactLink($(this)); });

})(jQuery);
