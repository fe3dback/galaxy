package vulkan

import "time"

// count of swap-chain images available for render into
const swapChainBuffersCount = 2

// GPU timeout for render. After this, app will be crashed
const swapChainTimeout = time.Second * 10

// How many frames can be failed continuously before crash
const maxPresetFails = 100
