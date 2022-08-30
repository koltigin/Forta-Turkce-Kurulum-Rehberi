# Forta Güncelleme

# Güncelleme Yöntemi-1
## Güncelleme
```shell
apt-get update
apt-get upgrade forta
```

## Service Dosyasını Düzenleme
```shell
nano /lib/systemd/system/forta.service   
```
Açılan dosyayı `ctrl k` ile sildikten sonra aşağıdaki kodda `SIFRENIZ` bölümünü kendinize göre düzenleyip yapıştırın ve dosyayı `ctrl x, y ve enter` ile kaydediniz. 

```shell
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

```

## Sistemi Yeniden Başlatma ve Durum Kontrolü
```shell
systemctl daemon-reload 
systemctl restart forta
systemctl status forta
```

# Güncelleme Yöntemi-2
```shell
systemctl stop forta
curl https://dist.forta.network/artifacts/forta -o $(which forta)
chmod +x /usr/bin/forta
forta version 
systemctl start forta
journalctl -fu forta
```
