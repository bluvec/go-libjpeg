CGO_CFLAGS="-D_GLIBCXX_USE_CXX11_ABI=0 -I${CONDA_PREFIX}/lib/python${PYTHON_VER}/site-packages/numpy/core/include -I${CONDA_PREFIX}/ -I${CONDA_PREFIX}/include"
CGO_CXXFLAGS="-I${CONDA_PREFIX}/lib/python${PYTHON_VER}/site-packages/numpy/core/include -I${CONDA_PREFIX}/lib/python${PYTHON_VER}/site-packages/torch/include -I${CONDA_PREFIX}/lib/python${PYTHON_VER}/site-packages/torch/include/torch/csrc/api/include"
CGO_LDFLAGS="-L${CONDA_PREFIX}/lib/python${PYTHON_VER}/site-packages/torch/lib -Wl,-rpath,${CONDA_PREFIX}/lib/python${PYTHON_VER}/site-packages/torch/lib -L${CONDA_PREFIX}/lib -Wl,-rpath,${CONDA_PREFIX}/lib"
export CGO_CFLAGS
export CGO_CXXFLAGS
export CGO_LDFLAGS
go build
