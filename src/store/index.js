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
      },
      comment: {
        isBtnPressed: false,
      },
      convert: {
        isBtnPressed: false,
      },
      btnSelection: {
        isFirstSelect: true,
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
    getBtnSelection(state) {
      return state.playgroundOptions.btnSelection;
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
    setbtnSelection(state, payload) {
      state.btnSelection.isFirstSelect = payload;
    },
  },
  actions: {},
  modules: {},
});
