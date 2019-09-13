#!/usr/bin/env bash
dir=$(pwd)
addons=(
base
web
webKanban
decimalPrecision
analytic
l10nGenericCoa
account
accountAsset
accountBudget
accountCheckPrinting
accountTaxCashBasis
accountVoucher
mail
mrp
mrpByProduct
procurement
product
project
projectIssue
purchase
purchaseRequisition
sale
saleMRP
saleMargin
saleOrderDates
saleStock
saleTeams
stock
stockAccount
stockLandedCost
)

echo  "Hexya Addons directory: $GOPATH/src/github.com/hexya-addons/"
mkdir -p "$GOPATH/src/github.com/hexya-addons/"
cd "$GOPATH/src/github.com/hexya-addons/"

for arrItem in ${addons[@]}
do
    if [ -d $arrItem ]; then
        echo "Checkout $arrItem ..."
        cd $arrItem && git pull && cd ../
    else
        echo "go get -u $arrItem"
        go get -u "github.com/hexya-addons/$arrItem"
    fi
done
cd $dir
