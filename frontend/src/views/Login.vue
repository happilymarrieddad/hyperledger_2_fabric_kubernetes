<template>
  <div class="login" style="margin-top:100px">
    <form>
        <div class="uk-margin">
            <input v-model="state.email" placeholder="Email" />
            <input v-model="state.password" placeholder="Password" />
        </div>
        <div uk-form-custom>
            <button type="button" @click="login">Login</button>
        </div>
    </form>
  </div>
</template>

<script>
import { reactive } from 'vue'
import { useStore } from "vuex";
import { useRouter } from 'vue-router';
export default {
    setup() {
        const store = useStore();
        const router = useRouter()
        const state = reactive({
            email: '',
            password: ''
        })
        const login = async () => {
            const [res, err] = await store.dispatch('login', {
                email: state.email,
                password: state.password
            })

            console.log(res);
            console.log(err);

            router.push('/');
        }
        return { 
            login,
            state
        }
    }
}
</script>