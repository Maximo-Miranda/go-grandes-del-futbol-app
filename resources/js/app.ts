import { createApp, h } from "vue";
import { createInertiaApp } from "@inertiajs/vue3";
import vuetify from "./plugins/vuetify";

import AppLayout from "./layouts/AppLayout.vue";
import AuthLayout from "./layouts/AuthLayout.vue";

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
