$(function () {
    $('form').on('submit', function (e) {
        e.preventDefault()

        var self = $(this)

        var formData = {
            name: self.find('[name="name"]').val(),
            gender: self.find('[name="gender"]').val(),
        }

        var url = self.attr('action')
        var method = self.attr('method')
        var payload = JSON.stringify(formData)

        $.ajax({
            url: url,
            type: method,
            contentType: 'application/json',
            data: payload,
            beforeSend: function(req) {
                var csrfToken = self.find('[name=csrf]').val()
                req.setRequestHeader("X-CSRF-Token", csrfToken)
            },
        }).then(function (res) {
            alert(res)
        }).catch(function (err) {
            alert('ERROR: ' + err.responseText)
            console.log('err', err)
        })
    })
})