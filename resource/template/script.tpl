<script>
const App = {
    components: {
        'picture-icon': ElementPlusIconsVue.Picture,
        'view-icon': ElementPlusIconsVue.View,
        'new-icon': ElementPlusIconsVue.Notification,
        'user-icon': ElementPlusIconsVue.User
    },
    data() {
        return {
            uploadData: {
                nsfw: false,
            },
            uploadResult: []
        };
    },
    mounted: function () {
        //
    },
    methods: {
        goToGitHub: function () {
            window.open("https://github.com/HunterXuan")
        },
        goToDockerHub: function () {
            window.open("https://hub.docker.com/r/hunterxuan")
        },
        goToBlog: function () {
            window.open("https://hunterx.xyz")
        },
        handleUploadSuccess: function (response, uploadFile, uploadFiles) {
            if (response.data) {
                this.uploadResult.push(response.data)
            }
        },
    }
};
const app = Vue.createApp(App)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
app.mount("#app")
</script>