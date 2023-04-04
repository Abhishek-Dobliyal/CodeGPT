import { createStore } from "vuex";

export default createStore({
  state: {
    editorInstance: null,
    editorContent: {
      text: "",
      language: "",
      highlightedText: "",
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
    getEditorInstance(state) {
      return state.editorInstance;
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
    setEditorInstance(state, payload) {
      state.editorInstance = payload;
    },
  },
  actions: {},
  modules: {},
});
