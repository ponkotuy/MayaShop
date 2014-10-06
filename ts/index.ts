/// <reference path="typings/jquery/jquery.d.ts" />
/// <reference path="vuejs.d.ts" />

$(() => {
    var vue = new Vue({
        el: "#login_form",
        data: {
            user: "",
            pass: ""
        },
        methods: {
            login: () => {
                var data = {name: this.user, passwd: this.pass}
                $.post('/api/v1/user/info', data)
            }
        }
    })
})
