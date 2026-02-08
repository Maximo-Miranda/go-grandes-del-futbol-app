import { reactive } from "vue";

const state = reactive({
  visible: false,
  text: "",
  color: "success",
});

let timer: ReturnType<typeof setTimeout> | null = null;

export function useToast() {
  const show = (msg: string, c = "success") => {
    if (timer) clearTimeout(timer);
    state.text = msg;
    state.color = c;
    state.visible = true;
    timer = setTimeout(() => {
      state.visible = false;
    }, 4000);
  };

  const close = () => {
    state.visible = false;
  };
  const success = (msg: string) => show(msg, "success");
  const error = (msg: string) => show(msg, "error");
  const warning = (msg: string) => show(msg, "warning");

  return { state, show, close, success, error, warning };
}
