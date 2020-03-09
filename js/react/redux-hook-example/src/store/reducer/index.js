const INIT_STATE={data:0}
export function countReducer(state = INIT_STATE, action) {
  switch (action.type) {
    case "ADD":
      return { data: state.data + 1 };
    case "SUB":
      return { data: state.data - 1 };
    default:
        return state;
  }
}