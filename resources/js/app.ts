import { createApp, h } from "vue";
import { createInertiaApp, router } from "@inertiajs/vue3";
import vuetify from "./plugins/vuetify";
import { useToast } from "./composables/useToast";

import AppLayout from "./layouts/AppLayout.vue";
import AuthLayout from "./layouts/AuthLayout.vue";

const toast = useToast();

router.on("invalid", (event) => {
  event.preventDefault();

  const response = event.detail.response;
  if (response.status === 429) {
    toast.warning("Demasiados intentos. Espera un momento.");
  } else if (response.status >= 500) {
    toast.error("Error del servidor. Intenta de nuevo.");
  } else {
    toast.error("Ocurrió un error inesperado.");
  }
});

router.on("exception", (event) => {
  event.preventDefault();
  toast.error("Error de conexión. Verifica tu internet.");
});

createInertiaApp({
  resolve: (name) => {
    const pages = import.meta.glob("./pages/**/*.vue", { eager: true });
    const page = pages[`./pages/${name}.vue`] as { default: any } | undefined;

    if (!page) {
      throw new Error(
        `Page not found: ${name}. Available: ${Object.keys(pages).join(", ")}`,
      );
    }

    const component = page.default;

    if (name.startsWith("auth/")) {
      component.layout = component.layout || AuthLayout;
    } else if (name !== "Welcome") {
      component.layout = component.layout || AppLayout;
    }

    return component;
  },
  setup({ el, App, props, plugin }) {
    createApp({ render: () => h(App, props) })
      .use(plugin)
      .use(vuetify)
      .mount(el);
  },
});
