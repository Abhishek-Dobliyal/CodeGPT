import { createStore } from "vuex";

export default createStore({
  state: {
    editorContent: {
      text: "",
      language: "",
    },
    playgroundOptions: {
      refactor: {
        isBtnPressed: false,
        splitCode: false,
        addTests: false,
      },
      prettify: {
        isBtnPressed: false,
        indentSpace: 2,
        useBestPractice: false,
      },
      comment: {
        isBtnPressed: false,
        commentEveryLine: false,
        brief: false,
      },
      convert: {
        isBtnPressed: false,
        to: "",
        start: -1,
        end: -1,
      },
    },
  },
  getters: {
    getEditorContent(state) {
      return state.editorContent;
    },
    getPlaygroundOptions(state) {
      return state.playgroundOptions;
    },
  },
  mutations: {
    setEditorAttrs(state, payload) {
      state.editorContent[payload.property] = payload.val;
    },
    setEditorText(state, payload) {
      state.editorContent.text = payload;
    },
    setPlaygroundOptions(state, payload) {
      state.playgroundOptions = payload;
    },
  },
  actions: {},
  modules: {},
});
