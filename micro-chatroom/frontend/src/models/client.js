export default {
  namespace: 'client',
  state: {
      pc: !/Android|webOS|iPhone|iPod|BlackBerry/i.test(navigator.userAgent),
  },
  reducers: {
    'change'(state, {
      pc,
    }) {
      return {
          pc,
      }
    },
  },
};
