module free5gclib

go 1.20

replace (
        stgutg => ./src/stgutg
        tglib => ./src/tglib
)

require (
        github.com/aead/cmac v0.0.0-20160719120800-7af84192f0b1
        github.com/antihax/optional v1.0.0
        github.com/antonfisher/nested-logrus-formatter v1.3.1
        github.com/calee0219/fatal v0.0.1
        github.com/dgrijalva/jwt-go v1.0.2
        github.com/evanphx/json-patch v0.5.2
        github.com/free5gc/nas v1.0.1
        github.com/free5gc/ngap v1.0.4
        github.com/free5gc/openapi v1.0.2
        github.com/ghedo/go.pkt v0.0.0-20200209120728-c97f47ad982f
        github.com/gin-gonic/gin v1.7.1
        github.com/google/uuid v1.3.0
        github.com/ishidawataru/sctp v0.0.0-20210707070123-9a39160e9062
        github.com/mitchellh/mapstructure v1.5.0
        github.com/pkg/errors v0.9.1
        github.com/sirupsen/logrus v1.9.0
        github.com/stretchr/testify v1.8.1
        go.mongodb.org/mongo-driver v1.11.2
        gopkg.in/yaml.v2 v2.4.0
        stgutg v0.0.0-00010101000000-000000000000
        tglib v0.0.0-00010101000000-000000000000
)

