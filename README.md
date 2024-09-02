# ImageGo

A free image-hosting service built with GoFrame. Demo: [iEndPot](https://i.endpot.com).

# Features

1. Built with GoFrame
2. All images are stored on s3-compatible storage provider (For example: B2 Cloud Storage, a High Performance Cloud Storage at 1/4 the Price)
3. Only hot images are cached (hot images are those viewed by visitors recently)
4. Dupe checking: When a duplicate image is uploaded, a connection would be created at the database and no disk space would be consumed.
5. NFSW tag support
6. and so on...

# How to use

## Docker

Docker is recommended to run ImageGo. You can use the following command to mount your config file and run it.

```shell
docker run -v ${your_config_file_path}:/app/manifest/config/config.yaml endpot/imagego
```

## Contributing

Thank you for considering contributing to this project! Feel free to raise your question, share your ideas or make a pull request.

# License

This project is open-sourced software licensed under the [MIT license](https://github.com/HunterXuan/ImageX/blob/master/LICENSE).
