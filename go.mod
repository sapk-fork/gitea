module code.gitea.io/gitea

go 1.12

require (
	gitea.com/macaron/binding v0.0.0-20190822013154-a5f53841ed2b
	gitea.com/macaron/cache v0.0.0-20190822004001-a6e7fee4ee76
	gitea.com/macaron/captcha v0.0.0-20190822015246-daa973478bae
	gitea.com/macaron/cors v0.0.0-20190821152825-7dcef4a17175
	gitea.com/macaron/csrf v0.0.0-20190822024205-3dc5a4474439
	gitea.com/macaron/i18n v0.0.0-20190822004228-474e714e2223
	gitea.com/macaron/inject v0.0.0-20190805023432-d4c86e31027a
	gitea.com/macaron/macaron v1.3.3-0.20190821202302-9646c0587edb
	gitea.com/macaron/session v0.0.0-20190821211443-122c47c5f705
	gitea.com/macaron/toolbox v0.0.0-20190822013122-05ff0fc766b7
	github.com/PuerkitoBio/goquery v0.0.0-20170324135448-ed7d758e9a34
	github.com/RoaringBitmap/roaring v0.4.7 // indirect
	github.com/andybalholm/cascadia v0.0.0-20161224141413-349dd0209470 // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/blevesearch/bleve v0.0.0-20190214220507-05d86ea8f6e3
	github.com/blevesearch/blevex v0.0.0-20180227211930-4b158bb555a3 // indirect
	github.com/blevesearch/go-porterstemmer v0.0.0-20141230013033-23a2c8e5cf1f // indirect
	github.com/blevesearch/segment v0.0.0-20160105220820-db70c57796cc // indirect
	github.com/boombuler/barcode v0.0.0-20161226211916-fe0f26ff6d26 // indirect
	github.com/chaseadamsio/goorgeous v0.0.0-20170901132237-098da33fde5f
	github.com/couchbase/vellum v0.0.0-20190111184608-e91b68ff3efe // indirect
	github.com/cznic/b v0.0.0-20181122101859-a26611c4d92d // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/cznic/strutil v0.0.0-20181122101858-275e90344537 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20190724012636-11b2859924c1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/emirpasic/gods v1.12.0
	github.com/etcd-io/bbolt v1.3.2 // indirect
	github.com/ethantkoenig/rupture v0.0.0-20180203182544-0a76f03a811a
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/facebookgo/ensure v0.0.0-20160127193407-b4ab57deab51 // indirect
	github.com/facebookgo/freeport v0.0.0-20150612182905-d4adf43b75b9 // indirect
	github.com/facebookgo/grace v0.0.0-20160926231715-5729e484473f
	github.com/facebookgo/httpdown v0.0.0-20160323221027-a3b1354551a2 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/stats v0.0.0-20151006221625-1b76add642e4 // indirect
	github.com/facebookgo/subset v0.0.0-20150612182917-8dac2c3c4870 // indirect
	github.com/gliderlabs/ssh v0.2.2
	github.com/glycerine/go-unsnap-stream v0.0.0-20180323001048-9f0cb55181dd // indirect
	github.com/glycerine/goconvey v0.0.0-20190315024820-982ee783a72e // indirect
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.4
	github.com/gogits/chardet v0.0.0-20150115103509-2404f7772561
	github.com/gogs/cron v0.0.0-20171120032916-9f6c956d3e14
	github.com/google/go-github/v24 v24.0.1
	github.com/gorilla/context v1.1.1
	github.com/issue9/assert v1.3.2 // indirect
	github.com/issue9/identicon v0.0.0-20160320065130-d36b54562f4c
	github.com/jaytaylor/html2text v0.0.0-20160923191438-8fb95d837f7d
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20170619183022-cd60e84ee657
	github.com/keybase/go-crypto v0.0.0-20170605145657-00ac4db533f6
	github.com/klauspost/compress v0.0.0-20161025140425-8df558b6cb6f
	github.com/klauspost/cpuid v0.0.0-20160302075316-09cded8978dc // indirect
	github.com/klauspost/crc32 v0.0.0-20161016154125-cb6bfca970f6 // indirect
	github.com/lafriks/xormstore v1.1.0
	github.com/lib/pq v1.2.0
	github.com/lunny/dingtalk_webhook v0.0.0-20171025031554-e3534c89ef96
	github.com/lunny/levelqueue v0.0.0-20190217115915-02b525a4418e
	github.com/markbates/goth v1.49.0
	github.com/mattn/go-isatty v0.0.7
	github.com/mattn/go-oci8 v0.0.0-20190320171441-14ba190cf52d // indirect
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/mcuadros/go-version v0.0.0-20190308113854-92cdf37c5b75
	github.com/microcosm-cc/bluemonday v0.0.0-20161012083705-f77f16ffc87a
	github.com/mschoch/smat v0.0.0-20160514031455-90eadee771ae // indirect
	github.com/msteinert/pam v0.0.0-20151204160544-02ccfbfaf0cc
	github.com/nfnt/resize v0.0.0-20160724205520-891127d8d1b5
	github.com/oliamb/cutter v0.2.2
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pquerna/otp v0.0.0-20160912161815-54653902c20e
	github.com/prometheus/client_golang v0.9.3
	github.com/remyoudompheng/bigfft v0.0.0-20190321074620-2f0d2b0e0001 // indirect
	github.com/russross/blackfriday v0.0.0-20180428102519-11635eb403ff
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sergi/go-diff v1.0.0
	github.com/shurcooL/httpfs v0.0.0-20190527155220-6a4d4a70508b // indirect
	github.com/shurcooL/sanitized_anchor_name v0.0.0-20160918041101-1dba4b3954bc // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/steveyen/gtreap v0.0.0-20150807155958-0abe01ef9be2 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/tecbot/gorocksdb v0.0.0-20181010114359-8752a9433481 // indirect
	github.com/tinylib/msgp v0.0.0-20180516164116-c8cf64dff200 // indirect
	github.com/tstranex/u2f v1.0.0
	github.com/unknwon/cae v0.0.0-20190822084630-55a0b64484a1
	github.com/unknwon/com v0.0.0-20190804042917-757f69c95f3e
	github.com/unknwon/i18n v0.0.0-20190805065654-5c6446a380b6
	github.com/unknwon/paginater v0.0.0-20151104151617-7748a72e0141
	github.com/urfave/cli v1.20.0
	github.com/willf/bitset v0.0.0-20180426185212-8ce1146b8621 // indirect
	github.com/yohcop/openid-go v0.0.0-20160914080427-2c050d2dae53
	go.etcd.io/bbolt v1.3.2 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421
	golang.org/x/sys v0.0.0-20190801041406-cbf593c0f2f3
	golang.org/x/text v0.3.2
	golang.org/x/tools v0.0.0-20190731214159-1e85ed8060aa // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20150924051756-4e86f4367175 // indirect
	gopkg.in/editorconfig/editorconfig-core-go.v1 v1.3.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.46.0
	gopkg.in/ldap.v3 v3.0.2
	gopkg.in/src-d/go-billy.v4 v4.3.2
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/stretchr/testify.v1 v1.2.2 // indirect
	gopkg.in/testfixtures.v2 v2.5.0
	mvdan.cc/xurls/v2 v2.0.0
	strk.kbt.io/projects/go/libravatar v0.0.0-20160628055650-5eed7bff870a
	xorm.io/builder v0.3.5
	xorm.io/core v0.7.0
)
