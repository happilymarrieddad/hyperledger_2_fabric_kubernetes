const state = {
    loading: false,
    list: [
        { id: 1, number: '1234' },
        { id: 2, number: '4321' }
    ]
};

const mutations = {
    addToList(state, order = {}) {
        state.list.push(order);
    },
    setList(state, list = []) {
        state.list = list;
    }
};
const getters = {
    getList(state) {
        return state.list || [];
    }
};

const actions = {
    getAll({ state }, { limit, offset } = {}) {
        state.loading = true;

        return new Promise((resolve) => {
            state.loading = false;
            return resolve([state.list, null]);
        })
    },
    createOne({ state, commit, dispatch }, { number }) {
        state.loading = true;

        const newOrder = { number };
        const lastExistingOrder = state.list.reverse()[0];
        newOrder.id = lastExistingOrder.id + 1;

        commit('addToList', newOrder);

        return dispatch('getAll');
    },
    destroyOne({ state, commit }, id = 0) {
        state.loading = true;

        commit('setList', state.list.filter(el => el.id !== id));

        return dispatch('getAll');
    },
    getOne({ state }, id = 0) {
        state.loading = true;

        return new Promise((resolve) => {
            return resolve([state.list.find(el => el.id === id), null]);
        })
    },
    updateOne({ state, dispatch }, { id, number }) {
        state.loading = true;

        const existingOrder = state.list.find(el => el.id === id);
        const existingOrderIndex = state.list.findIndex(el => el.id === id);

        return new Promise((resolve) => {
            if (!existingOrder) {
                return resolve(null, new Error("unable to find existing order"))
            }

            existingOrder.number = number;
            state.list[existingOrderIndex] = existingOrder

            return dispatch('getAll');
        });
    }
}

export default {
    namespaced: true,
    name: 'orders',
    state,
    mutations,
    actions,
    getters
}
