<template>
    <div class="container my-5">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Register</div>
                    <div class="card-body">
                        <div v-if="error" class="alert alert-danger">{{ error }}</div>

                        <form action="#" @submit.prevent="submit">
                            <div class="form-group row mb-3">
                                <label for="name" class="col-md-4 col-form-label text-md-right">Name</label>

                                <div class="col-md-6">
                                    <input
                                        id="name"
                                        type="name"
                                        class="form-control"
                                        name="name"
                                        value
                                        required
                                        autofocus
                                        v-model="form.name"
                                    />
                                </div>
                            </div>

                            <div class="form-group row mb-3">
                                <label for="email" class="col-md-4 col-form-label text-md-right">Email</label>

                                <div class="col-md-6">
                                    <input
                                        id="email"
                                        type="email"
                                        class="form-control"
                                        name="email"
                                        value
                                        required
                                        autofocus
                                        v-model="form.email"
                                    />
                                </div>
                            </div>

                            <div class="form-group row mb-3">
                                <label for="password" class="col-md-4 col-form-label text-md-right">Password</label>

                                <div class="col-md-6">
                                    <input
                                        id="password"
                                        type="password"
                                        class="form-control"
                                        name="password"
                                        required
                                        v-model="form.password"
                                    />
                                </div>
                            </div>

                            <div class="form-group row mb-0">
                                <div class="col-md-8 offset-md-4">
                                    <button type="submit" class="btn btn-primary">Register</button>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>


<script>
import { mapActions, mapGetters } from 'vuex'

export default {
    name: 'Register',
    data: function() {
        return {
            form: {
                name: "",
                email: "",
                password: ""
            },
            error: null
        }
    },
    watch: {
        loggedIn: function(value, _old) {
            if (value) {
                this.$router.replace({ name: 'Lobby' })
            }
        }
    },
    computed: {
        ...mapGetters([
            'loggedIn'
        ])
    },
    methods: {
        ...mapActions([
            'register'
        ]),
        submit: async function() {
            try {
                await this.register({ 
                    username: this.form.name, 
                    email: this.form.email, 
                    password: this.form.password 
                })

                this.$router.replace({ name: 'Auth' })
            } catch(err) {
                this.error = err.message
            }
        }
    },
    mounted: function() {
        if (this.loggedIn) {
            this.$router.replace({ name: 'Lobby' })
        }
    }
}
</script>