import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import "vuetify/styles";
import "@mdi/font/css/materialdesignicons.css";

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: "light",
    themes: {
      light: {
        dark: false,
        colors: {
          primary: "#1B5E20",
          "primary-darken-1": "#145218",
          secondary: "#4CAF50",
          "secondary-darken-1": "#388E3C",
          accent: "#FFC107",
          error: "#F44336",
          info: "#2196F3",
          success: "#66BB6A",
          warning: "#FB8C00",
          background: "#F1F8E9",
          surface: "#FFFFFF",
          "surface-variant": "#E8F5E9",
          "on-surface-variant": "#33691E",
        },
      },
    },
  },
  defaults: {
    VBtn: {
      variant: "flat",
      rounded: "lg",
    },
    VCard: {
      rounded: "lg",
    },
    VTextField: {
      variant: "outlined",
      density: "comfortable",
      color: "primary",
    },
    VSelect: {
      variant: "outlined",
      density: "comfortable",
      color: "primary",
    },
    VTextarea: {
      variant: "outlined",
      density: "comfortable",
      color: "primary",
    },
  },
});

export default vuetify;
