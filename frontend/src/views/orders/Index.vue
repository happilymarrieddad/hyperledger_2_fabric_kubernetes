<template>
<div class="orders">
    <ul class="uk-breadcrumb">
        <li><span>orders</span></li>
        <li>
            <router-link to="/orders/create"><span uk-icon="plus-circle"></span></router-link>
        </li>
    </ul>

    <table class="uk-table">
        <thead>
            <tr>
                <th>ID</th>
                <th>Number</th>
                <th></th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="order in state.orders" v-bind:key="order.id">
                <td v-html="order.id"></td>
                <td v-html="order.number"></td>
                <td>
                    <router-link :to="'/orders/edit/'+order.id"><span class="icon-style" uk-icon="settings"></span></router-link>
                    <span class="icon-style" @click="destroy(order.id)" uk-icon="trash"></span>
                </td>
            </tr>
        </tbody>
    </table>
</div>
</template>

<script>
import {
    reactive,
    computed,
    onMounted
} from 'vue'
import {
    useStore
} from "vuex";
import uikit from 'uikit';

export default {
    name: 'orders',
    setup() {
        const store = useStore();

        const state = reactive({
            orders: computed(() => store.state.orders.list)
        })

        const fetchData = () => {
            store.dispatch('orders/getAll')
        }

        const destroy = (id = 0) => {
            uikit.modal.confirm('Are you sure?').then(async () => {
                await store.dispatch('orders/destroyOne', id)

                fetchData()
            });
        }

        onMounted(fetchData);

        return {
            destroy,
            state
        }
    }
}
</script>