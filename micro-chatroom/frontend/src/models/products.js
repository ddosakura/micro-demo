export default {
  namespace: 'products',
  state: [{
    name: 'demo',
    id: 0,
  }, ],
  reducers: {
    'delete'(state, {
      payload: id
    }) {
      return state.filter(item => item.id !== id);
    },
    'fresh'(state, {
      payload: data
    }) {
      return [...state, ...data] 
    },
  },
};
