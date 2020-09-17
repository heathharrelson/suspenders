<template>
  <nav class="bg-gray-800">
    <div class="max-w-7xl mx-auto px-4 sm:px-4 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <router-link to="/" class="font-semibold text-white">
              Suspenders
            </router-link>
          </div>
          <div class="hidden sm:block sm:ml-6">
            <div class="flex">
              <router-link
                to="/"
                class="px-3 py-2 rounded-md text-sm font-medium leading-5 text-gray-300 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out"
                exact-active-class="text-white bg-gray-900"
                exact
              >
                Deployments
              </router-link>
              <router-link
                to="/about"
                class="px-3 py-2 rounded-md text-sm font-medium leading-5 text-gray-300 focus:outline-none focus:text-white focus:bg-gray-700 transition duration-150 ease-in-out"
                active-class="text-white bg-gray-900"
              >
                About
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>

  <main class="flex flex-col flex-grow">
    <Suspense>
      <template #default>
        <router-view />
      </template>
      <template #fallback>
        <p class="m-auto text-gray-500">
          Loading
          <svg
            viewBox="0 0 20 20"
            fill="currentColor"
            class="inline-block h-4 animate-spin"
          >
            <path
              fill-rule="evenodd"
              d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z"
              clip-rule="evenodd"
            />
          </svg>
        </p>
      </template>
    </Suspense>
  </main>
</template>
