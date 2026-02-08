<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { Head, usePage, router, Link } from "@inertiajs/vue3";
import { useToast } from "../composables/useToast";

interface PageProps extends Record<string, unknown> {
    flash?: { success?: string; error?: string; warning?: string };
    auth?: { user: { id: number; name: string; email: string; role: string } };
}

const page = usePage<PageProps>();
const drawer = ref(true);
const rail = ref(false);
const mobileDrawer = ref(false);
const flash = computed(() => page.props.flash);
const authUser = computed(() => page.props.auth?.user);
const toast = useToast();

const isMobile = computed(() => {
    if (typeof window === "undefined") return false;
    return window.innerWidth < 960;
});

onMounted(() => {
    const savedRail = localStorage.getItem("drawer-rail");
    if (savedRail !== null) rail.value = savedRail === "true";
    if (flash.value?.success) toast.success(flash.value.success);
    else if (flash.value?.error) toast.error(flash.value.error);
    else if (flash.value?.warning) toast.warning(flash.value.warning);
});

const toggleRail = () => {
    rail.value = !rail.value;
    localStorage.setItem("drawer-rail", String(rail.value));
};

const logout = () => router.post("/auth/logout");

const navItems = [
    { title: "Dashboard", icon: "mdi-view-dashboard", to: "/dashboard" },
    { title: "Torneos", icon: "mdi-trophy", to: "/tournaments" },
    { title: "Equipos", icon: "mdi-shield-half-full", to: "/teams" },
    { title: "Jugadores", icon: "mdi-account-group", to: "/players" },
    { title: "Sedes", icon: "mdi-map-marker", to: "/venues" },
    { title: "Partidos", icon: "mdi-soccer", to: "/matches" },
    {
        title: "Clasificación",
        icon: "mdi-format-list-numbered",
        to: "/standings",
    },
];
</script>

<template>
    <Head><title>Grandes del Fútbol</title></Head>
    <v-app>
        <v-app-bar color="primary" density="comfortable" flat>
            <v-app-bar-nav-icon
                v-if="isMobile"
                @click="mobileDrawer = !mobileDrawer"
            />
            <v-app-bar-nav-icon v-else @click="toggleRail" />
            <v-app-bar-title>
                <Link
                    href="/dashboard"
                    class="text-white text-decoration-none d-flex align-center"
                >
                    <span class="mr-2">⚽</span>
                    <span class="font-weight-medium">Grandes del Fútbol</span>
                </Link>
            </v-app-bar-title>
            <v-spacer />
            <v-menu offset-y>
                <template #activator="{ props }">
                    <v-btn v-bind="props" variant="text" class="text-white">
                        <v-avatar
                            size="32"
                            color="primary-darken-1"
                            class="mr-2"
                        >
                            <span class="text-body-2">{{
                                authUser?.name?.charAt(0) || "U"
                            }}</span>
                        </v-avatar>
                        <span class="d-none d-sm-inline">{{
                            authUser?.name || "Usuario"
                        }}</span>
                        <v-icon end>mdi-chevron-down</v-icon>
                    </v-btn>
                </template>
                <v-list density="compact">
                    <v-list-item prepend-icon="mdi-logout" @click="logout">
                        <v-list-item-title>Cerrar Sesión</v-list-item-title>
                    </v-list-item>
                </v-list>
            </v-menu>
        </v-app-bar>

        <v-navigation-drawer
            v-if="!isMobile"
            v-model="drawer"
            :rail="rail"
            permanent
        >
            <v-list density="compact" nav>
                <v-list-item
                    v-for="item in navItems"
                    :key="item.to"
                    :prepend-icon="item.icon"
                    :title="item.title"
                    :href="item.to"
                    rounded="lg"
                    color="primary"
                />
            </v-list>
            <template #append>
                <v-list density="compact" nav>
                    <v-list-item
                        prepend-icon="mdi-logout"
                        title="Cerrar Sesión"
                        @click="logout"
                        rounded="lg"
                    />
                </v-list>
            </template>
        </v-navigation-drawer>

        <v-navigation-drawer v-if="isMobile" v-model="mobileDrawer" temporary>
            <v-list density="compact" nav>
                <v-list-item
                    v-for="item in navItems"
                    :key="item.to"
                    :prepend-icon="item.icon"
                    :title="item.title"
                    :href="item.to"
                    rounded="lg"
                    color="primary"
                    @click="mobileDrawer = false"
                />
            </v-list>
            <template #append>
                <v-list density="compact" nav>
                    <v-list-item
                        prepend-icon="mdi-logout"
                        title="Cerrar Sesión"
                        @click="logout"
                        rounded="lg"
                    />
                </v-list>
            </template>
        </v-navigation-drawer>

        <v-main class="bg-grey-lighten-4">
            <v-container fluid class="pa-4 pa-md-6">
                <slot />
            </v-container>
        </v-main>

        <v-snackbar
            v-model="toast.state.visible"
            :color="toast.state.color"
            :timeout="4000"
            location="top center"
        >
            {{ toast.state.text }}
            <template #actions>
                <v-btn variant="text" @click="toast.close()">Cerrar</v-btn>
            </template>
        </v-snackbar>
    </v-app>
</template>
