<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>ImageGo</title>
    <link rel="shortcut icon" href="/favicon.ico">
    <link rel="stylesheet" href="//unpkg.com/element-plus/dist/index.css" />
    <link rel="stylesheet" href="//unpkg.com/prismjs/themes/prism.css" rel="stylesheet" />
    <script src="//unpkg.com/vue@3"></script>
    <script src="//unpkg.com/element-plus"></script>
    <script src="//unpkg.com/@element-plus/icons-vue"></script>
    <script src="//unpkg.com/prismjs/prism.js"></script>
    {{ include "style.tpl" }}
</head>
<body style="margin: 0">
    <div id="app" v-loading="loading">
        <el-row class="landing-row" justify="center" align="middle">
            <el-col :span="16">
                <el-card style="max-width: 100%; margin-top: 32px; margin-bottom: 32px;">
                    <template #header>
                        <div class="card-header">
                            <div style="display: flex; align-items: center;">
                                <div><el-image style="height: 48px;" src="/assets/img/logo.png" /></div>
                                <span class="el-divider el-divider--vertical" style="height: 48px;"></span>
                                <div>
                                    <div>
                                        <el-tooltip
                                            effect="dark"
                                            content="Image Count"
                                            placement="top-start"
                                            >
                                            <el-button class="stats-tag" type="primary" size="small" plain>
                                                <el-icon><picture-icon /></el-icon>
                                                <span>{{ .imageCount }}</span>
                                            </el-button>
                                        </el-tooltip>
                                        <el-tooltip
                                            effect="dark"
                                            content="View Count"
                                            placement="top-start"
                                            >
                                            <el-button class="stats-tag" type="success" size="small" plain>
                                                <el-icon><view-icon /></el-icon>
                                                <span>{{ .imageView }}</span>
                                            </el-button>
                                        </el-tooltip>
                                        <el-tooltip
                                            effect="dark"
                                            content="New Image"
                                            placement="top-start"
                                            >
                                            <el-button class="stats-tag" type="info" size="small" plain>
                                                <el-icon><new-icon /></el-icon>
                                                <span>{{ .newImageCount }}</span>
                                            </el-button>
                                        </el-tooltip>
                                        <el-tooltip
                                            effect="dark"
                                            content="Uploader Count"
                                            placement="top-start"
                                            >
                                            <el-button class="stats-tag" type="warning" size="small" plain>
                                                <el-icon><user-icon /></el-icon>
                                                <span>{{ .uploaderCount }}</span>
                                            </el-button>
                                        </el-tooltip>
                                    </div>
                                    <el-text size="large">A simple, high-performance and cloud-ready image-hosting service written in golang, based on the GoFrame framework.</el-text>
                                </div>
                            </div>
                        </div>
                    </template>
                    <el-upload
                        drag
                        action="/api/upload"
                        multiple
                        name="image"
                        list-type="picture"
                        :data="uploadData"
                        :on-success="handleUploadSuccess"
                        >
                        <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                        <div class="el-upload__text">
                          Drop file here or <em>click to upload</em>
                        </div>
                        <template #tip>
                          <div class="el-upload__tip">
                            jpg/png/gif files with a size less than 5mb
                          </div>
                        </template>
                    </el-upload>
                    <div>
                        <el-switch
                            v-model="uploadData.nsfw"
                            class="mb-2"
                            active-text="NSFW"
                            inactive-text="SFW"
                        />
                    </div>
                    <div>
                        <el-tabs model-value="url" type="card">
                            <el-tab-pane label="URL" name="url">
                                <pre v-show="uploadResult.length"><code class="language-markup"><template v-for="item in uploadResult">{{ "{{ item.link }}\n" }}</template></code></pre>
                            </el-tab-pane>
                            <el-tab-pane label="HTML" name="html">
                                <pre v-show="uploadResult.length"><code class="language-markup"><template v-for="item in uploadResult">{{ "&lt;img src=\"{{ item.link }}\" alt=\"{{ item.name }}\" title=\"{{ item.name }}\" /&gt;\n" }}</template></code></pre>
                            </el-tab-pane>
                            <el-tab-pane label="BBCode" name="bbcode">
                                <pre v-show="uploadResult.length"><code class="language-markup"><template v-for="item in uploadResult">{{ "[img]{{ item.link }}[/img]\n" }}</template></code></pre>
                            </el-tab-pane>
                            <el-tab-pane label="Markdown" name="markdown">
                                <pre v-show="uploadResult.length"><code class="language-markup"><template v-for="item in uploadResult">{{ "![{{ item.name }}]({{ item.link }})\n" }}</template></code></pre>
                            </el-tab-pane>
                            <el-tab-pane label="Markdown with Link" name="markdown_with_link">
                                <pre v-show="uploadResult.length"><code class="language-markup"><template v-for="item in uploadResult">{{ "[![{{ item.name }}]({{ item.link }})]({{ item.link }})\n" }}</template></code></pre>
                            </el-tab-pane>
                            <el-tab-pane label="Delete Link" name="delete_link">
                                <pre v-show="uploadResult.length"><code class="language-markup"><template v-for="item in uploadResult">{{ "{{ item.delete_link }}\n" }}</template></code></pre>
                            </el-tab-pane>
                        </el-tabs>
                    </div>
                </el-card>
            </el-col>
        </el-row>

        <div><el-divider /></div>

        <div>
            <div class="recent-title">
                <h1>Recent Pictures</h1>
            </div>
        </div>

        <el-row class="recent-row" :gutter="20">
            {{ range $index, $elem := .recentImageList }}
                <el-col :span="4" style="padding-top: 16px;">
                    <el-card style="max-width: 100%" shadow="always">
                        <template #header><span><b>{{ $elem.Name }}</b></span></template>
                        <el-image
                          fit="contain"
                          src="/image/{{ $elem.Code }}/{{ $elem.Name }}"
                          :preview-src-list="['/image/{{ $elem.Code }}/{{ $elem.Name }}']"
                          :hide-on-click-modal="true"
                          style="width: 100%; height: 50vh;"
                          lazy
                        />
                      </el-card>
                </el-col>
            {{ end }}
        </el-row>

        <el-footer class="copyright-footer">
            <p>Copyright © {{ date "Y" }}. All rights reserved.</p>
        </el-footer>
    </div>

    {{ include "script.tpl" }}
</body>
</html>
