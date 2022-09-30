# Forta Network Türkçe Scan Node Kurulum Rehberi
![image](https://user-images.githubusercontent.com/102043225/179364468-8f51a9ca-c24b-449a-9588-d66104368089.png)
![Build](https://github.com/forta-network/forta-node/actions/workflows/release-codedeploy-dev.yml/badge.svg)

## Bilgi
Forta Network kurulumu yapabilmeniz için node'unuza 500 Forta token stake etmeniz gerekmektedir. Eylül ayından sonra stake miktarı artacak ve 2500 Forta token stake edenler node çalıştırabilecektir.

Forta node taşımak için [buraya](https://github.com/koltigin/Forta-Turkce-Kurulum-Rehberi/blob/master/Forta-Node-Tasima-Rehberi.md) bakabilirsiniz.

Forta node güncellemek için [buraya](https://github.com/koltigin/Forta-Turkce-Kurulum-Rehberi/blob/master/Forta-Guncelleme.md) bakabilirsiniz.

## Sistem Gereksinimleri
* 4CPU+ cores
* 16GB RAM
* Docker v20.10+
* 100GB SSD 

## Sistemi Güncelleme
```shell
sudo apt update && sudo apt upgrade -y
```

## Gerekli Kütüphanelerin Kurulması
```shell
sudo apt install ca-certificates curl gnupg lsb-release git htop liblz4-tool screen -y < "/dev/null"
```
## Docker Kurulumu
```shell
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
apt-get update
apt-get install docker-ce docker-ce-cli containerd.io
docker version
```
## Docker Daemon Dosyasının Oluşturulması
```shell
sudo tee /etc/docker/daemon.json > /dev/null <<EOF
{
   "default-address-pools": [
        {
            "base":"172.17.0.0/12",
            "size":16
        },
        {
            "base":"192.168.0.0/16",
            "size":20
        },
        {
            "base":"10.99.0.0/16",
            "size":24
        }
    ]
}
EOF
```

## Docker'ı Başlatma
```shell
systemctl restart docker
```

## Forta Kurulumu

`SIFRENIZ` bölümüne şifrenizi giriniz. Özel karakter kullanmayınız. Özel karakter kullanımında sorunlar oluşabiliyor.
```shell
sudo curl https://dist.forta.network/pgp.public -o /usr/share/keyrings/forta-keyring.asc -s
echo 'deb [signed-by=/usr/share/keyrings/forta-keyring.asc] https://dist.forta.network/repositories/apt stable main' | sudo tee -a /etc/apt/sources.list.d/forta.list
apt-get update
apt-get install forta
forta init --passphrase SIFRENIZ
```
Yukarıdaki kodların çıktısında aşağıdaki gibi Scanner adresinizi göreceksiniz. Bu adrese EVM cüzdanınızadan 0,1 MATIC gönderdikten sonra işlemlere devam ediyoruz.
```shell
Scanner address: 0xAAA8C491232cB65a65FBf7F36b71220B3E695AAA

Successfully initialized at /root/.forta
```

## Alchemy Hesap Oluşturma

[Alchemy](https://alchemy.com/?r=zc3NjI5NzM1NzMxN) adresine giderek bir hesap oluşturuyoruz. Burada `Create App` bölümünden `Polygon Mainnet App` oluşturuyoruz. Burada `View Key` bölümünden `https` ile başlayan linkimizi alıyoruz ve kurulum sırasında Alchemy linki geçen yerde kullanmak üzere bir txt dosyasına kaydediyoruz.
* Eğer node'unuzu başka bir ağda çalıştırmak isterseniz [buradan](https://docs.forta.network/en/latest/scanner-quickstart/) örnekleri görebilirsiniz. Çalıştırmak istediğiniz ağa göre alchemy app oluşturmanız gerekir.
 
## Yapılandırma Dosyası Oluşturma
Hangi ağda çalıştırmak istiyorsanız o ağın aşağıdaki yapılandırmalarını yapmanız gerekmektedir.

### Polygon:
```shell
rm /root/.forta/config.yml
sudo tee /root/.forta/config.yml > /dev/null <<EOF
chainId: 137

scan:
  jsonRpc:
    url: ALCHEMY_LINKINIZ

trace:
  enabled: false
EOF
```

### BSC:
```shell
rm /root/.forta/config.yml
sudo tee /root/.forta/config.yml > /dev/null <<EOF
chainId: 56

scan:
  jsonRpc:
    url: ALCHEMY_LINKINIZ

trace:
  enabled: false
EOF
```
### Ethereum:
```shell
rm /root/.forta/config.yml
sudo tee /root/.forta/config.yml > /dev/null <<EOF
chainId: 1

scan:
  jsonRpc:
    url: ALCHEMY_LINKINIZ

trace:
  enabled: false
EOF
```
### Avalanche:
```shell
rm /root/.forta/config.yml
sudo tee /root/.forta/config.yml > /dev/null <<EOF
chainId: 43114

scan:
  jsonRpc:
    url: ALCHEMY_LINKINIZ

trace:
  enabled: false
EOF
```
### Arbitrum:
```shell
rm /root/.forta/config.yml
sudo tee /root/.forta/config.yml > /dev/null <<EOF
chainId: 42161

scan:
  jsonRpc:
    url: ALCHEMY_LINKINIZ

trace:
  enabled: false
EOF
```
### Optimism:
```shell
rm /root/.forta/config.yml
sudo tee /root/.forta/config.yml > /dev/null <<EOF
chainId: 10

scan:
  jsonRpc:
    url: ALCHEMY_LINKINIZ

trace:
  enabled: false
EOF
```
### Fantom:
```shell
rm /root/.forta/config.yml
sudo tee /root/.forta/config.yml > /dev/null <<EOF
chainId: 250

scan:
  jsonRpc:
    url: ALCHEMY_LINKINIZ

trace:
  enabled: false
EOF
```

## Cüzdanı Kaydetme
`EVM_ADRESINIZ` bölümünde EVM cüzdan adresinizi ve `SIFRENIZ` yerine de yukarıda Forta kurulumunda belirlediğiniz şifreyi giriyoruz.
```shell
forta register --owner-address EVM_ADRESINIZ --passphrase SIFRENIZ
```

## Servis Dosyası Oluşturma
`SIFRENIZ` yerine yukarıda Forta kurulumunda belirlediğiniz şifreyi giriyoruz.
```shell
sudo tee /lib/systemd/system/forta.service > /dev/null <<EOF
[Unit]
Description=Forta
After=network-online.target
Wants=network-online.target systemd-networkd-wait-online.service

StartLimitIntervalSec=500
StartLimitBurst=5

[Service]
Environment="FORTA_DIR=/root/.forta/"
Environment="FORTA_PASSPHRASE=SIFRENIZ"
Restart=on-failure
RestartSec=15s

ExecStart=/usr/bin/forta run

[Install]
WantedBy=multi-user.target
EOF
```

## Forta Node Aktif Etme ve Çalıştırma
```shell
systemctl enable forta
systemctl daemon-reload
systemctl start forta
```

## Node Durumu
Node'unuzun SLA değerini kontrol etmek için bu adresin `https://api.forta.network/stats/sla/scanner/0x....` sonuna scanner adresinizi yazarak kontrol edebilirsiniz.

## Faydalı Komutlar

### Node Durumunu Görme
```shell
forta status
```

### Logları Görüntüleme

Çalışan konteynerleri görmek için;
```shell
docker ps
```
Bir konteynerı görmek için;
```shell
docker logs -f <container_id>
```

## Node Silme
```shell
systemctl stop forta
systemctl disable forta
rm /lib/systemd/system/forta.service -rf
rm $HOME/.forta -rf
```
