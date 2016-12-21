#!/bin/bash

set -e

echo "Generating email mocks"

mockery -name=Sender -case=underscore -inpkg
mockery -name=SenderSESAPI -case=underscore -inpkg

echo "Generating email mocks finished"
