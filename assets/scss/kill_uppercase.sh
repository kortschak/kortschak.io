#!/bin/bash

# This grabs all the scss definitions in Academic that contain an uppercase
# text-transform and writes them as a single custom.scss file with the uc
# transforms removed.

sed \
	-e 's/text-transform: uppercase;/text-transform: none;/g' \
	$(grep -lR 'text-transform: uppercase;' ../../themes/academic/assets/scss) \
> custom.scss
