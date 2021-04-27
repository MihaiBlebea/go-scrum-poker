<template>
    <div class="container my-5">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Login</div>
                    <div class="card-body">
                        <div v-if="error" class="alert alert-danger">{{ error }}</div>
                        <form action="#" @submit.prevent="submit">
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
                                    <button type="submit" class="btn btn-primary">Login</button>
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
import { mapActions } from 'vuex'

export default {
    data: function() {
        return {
            form: {
                email: "",
                password: ""
            },
            error: null
        }
    },
    methods: {
        ...mapActions([
            'login'
        ]),
        submit: async function() {
            try {
                let result = await this.login({ 
                    email: this.form.email, 
                    password: this.form.password 
                })
                
                console.log(result)
                this.$router.replace({ name: 'Lobby' })
                
            } catch(err) {
                console.error(err)
                this.error = err.message
            }
        }
    }
}
</script>